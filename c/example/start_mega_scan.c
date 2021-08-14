#include <gvm/util/mqtt.h>
#include <signal.h>
#include <stdio.h>
#include <stdlib.h>
#define TOPIC "#"
#include <eulabeia/client.h>
#define COMPLEX_START_SCAN_ID "mega_scan_666"
#define OPENVAS_SENSOR "localhorst"

struct EulabeiaClient *ec;

void signalhandler(int signum)
{
	printf("Caught %d, reset mqtt\n", signum);
	eulabeia_destroy(ec);
	exit(0);
}

static struct EulabeiaScan *create_mega_scan()
{
	struct EulabeiaScan *scan;
	struct EulabeiaTarget *target;
	struct EulabeiaHosts *t_hosts;
	struct EulabeiaPlugins *t_plugins;
	struct EulabeiaPorts *t_ports;
	scan = calloc(1, sizeof(struct EulabeiaScan));
	scan->id = COMPLEX_START_SCAN_ID;
	scan->temporary = 1;
	target = calloc(1, sizeof(struct EulabeiaTarget));
	target->alive = 1;
	target->sensor = OPENVAS_SENSOR;
	t_hosts = calloc(1, sizeof(struct EulabeiaHosts));
	t_hosts->hosts = calloc(2, sizeof(struct EulabeiaHost));
	t_hosts->hosts[0].address = "localhost";
	t_hosts->len = 1;
	t_hosts->cap = 2;
	target->hosts = t_hosts;
	t_plugins = calloc(1, sizeof(struct EulabeiaPlugins));
	t_plugins->cap = 1;
	t_plugins->len = 1;
	t_plugins->plugins = calloc(1, sizeof(struct EulabeiaPlugins));
	t_plugins->plugins[0].oid = "1.3.6.1.4.1.25623.1.0.90022";
	target->plugins = t_plugins;
	t_ports = calloc(1, sizeof(struct EulabeiaPorts));
	t_ports->cap = 1;
	t_ports->len = 1;
	t_ports->ports = calloc(1, sizeof(struct EulabeiaPorts));
	t_ports->ports[0].port = "22";
	target->ports = t_ports;
	scan->target = target;
	return scan;
}

static void free_mega_scan(struct EulabeiaScan *scan)
{
	free(scan->target->hosts->hosts);
	free(scan->target->hosts);
	free(scan->target->plugins->plugins);
	free(scan->target->plugins);
	free(scan->target->ports->ports);
	free(scan->target->ports);
	free(scan->target);
	free(scan);
}

static struct EulabeiaScan *start_mega_scan(struct EulabeiaScanProgress *sp)
{
	int rc;
	struct EulabeiaScan *scan;
	scan = create_mega_scan();

	if ((rc = eulabeia_start_scan(ec, scan, NULL)) != 0) {
		printf("[%d] unable to start scan %s\n", rc, scan->id);
		free_mega_scan(scan);
		eulabeia_destroy(ec);
		exit(rc);
	}
	return scan;
}

int main()
{
	struct EulabeiaScanProgress msp;
	struct EulabeiaScan *mega_scan;
	char *topic, *payload;
	int topic_len, payload_len;
	int rc;

	signal(SIGINT, signalhandler);
	if ((ec = eulabeia_initialize("localhost:9138", NULL)) == NULL) {
		printf("init returned NULL, quitting\n");
		rc = -1;
		goto exit;
	}

	mega_scan = start_mega_scan(&msp);
	while (!eulabeia_scan_finished(&msp)) {
		if ((rc = mqtt_retrieve_message(
			 &topic, &topic_len, &payload, &payload_len)) == -1) {
			printf("unable to retrieve message, quitting\n");
			goto exit;
		}
		if (rc == 0) {
			if ((rc = eulabeia_scan_progress(
				 payload, mega_scan->id, &msp)) == 0) {
				printf("[scan_id:%s][status:%d] %s\n",
				       mega_scan->id,
				       msp.status,
				       payload);
			} else {
				printf("payload:\n%s\n", payload);
				rc = 0;
			}
		} else {
			rc = 0;
		}
		if (payload != NULL)
			free(payload);
		if (topic != NULL)
			free(topic);
	}
exit:
	free_mega_scan(mega_scan);

	eulabeia_destroy(ec);
	return rc;
}
