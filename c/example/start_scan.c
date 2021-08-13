#include <gvm/util/mqtt.h>
#include <signal.h>
#include <stdio.h>
#include <stdlib.h>
#define TOPIC "#"
#include <eulabeia/client.h>
#define COMPLEX_START_SCAN_ID "classic_scan_1"
#define OPENVAS_SENSOR "localhorst"
#define TARGET_ID "example_target_1234"
#define GROUP_ID "example_group_1"

struct EulabeiaClient *ec;

void signalhandler(int signum)
{
	printf("Caught %d, reset mqtt\n", signum);
	eulabeia_destroy(ec);
	exit(0);
}

static struct EulabeiaTarget *example_target(){
	struct EulabeiaTarget *target;
	target = calloc(1, sizeof(struct EulabeiaTarget));
	target->id = TARGET_ID;
	target->alive = 1;
	target->sensor = OPENVAS_SENSOR;
	struct EulabeiaHosts *t_hosts;
	struct EulabeiaPlugins *t_plugins;
	struct EulabeiaPorts *t_ports;
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
	t_plugins->plugins[0].oid = "";
	target->plugins = t_plugins;
	t_ports = calloc(1, sizeof(struct EulabeiaPorts));
	t_ports->cap = 1;
	t_ports->len = 1;
	t_ports->ports = calloc(1, sizeof(struct EulabeiaPorts));
	t_ports->ports[0].port = "tcp://22";
	target->ports = t_ports;
	return target;
}

static struct EulabeiaScan *example_scan()
{
	struct EulabeiaScan *scan;
	scan = calloc(1, sizeof(struct EulabeiaScan));
	scan->id = COMPLEX_START_SCAN_ID;
	scan->target_id = TARGET_ID;
	return scan;
}

static void free_example_target(struct EulabeiaTarget *target)
{
	free(target->hosts->hosts);
	free(target->hosts);
	free(target->plugins->plugins);
	free(target->plugins);
	free(target->ports->ports);
	free(target->ports);
	free(target);
}

static void free_example_scan(struct EulabeiaScan *scan)
{
	free(scan);
}


static int check_for_modify_progress(struct EulabeiaCRUDProgress *progress, char *id){
	int rc;
	char *payload, *topic;
	int payload_len, topic_len;
	while (progress->status != EULABEIA_CRUD_SUCCESS 
			&& progress->status != EULABEIA_CRUD_FAILED){
		if ((rc = ec->retrieve(
			 &topic, &topic_len, &payload, &payload_len, NULL)) == -1) {
			printf("unable to retrieve message, quitting\n");
			rc = -1;
		}
		if (rc == 0) {
			if ((rc = eulabeia_modify_progress(
				 payload, id, progress)) == 0) {
				printf("[id:%s][status:%d] %s\n",
				       id,
				       progress->status,
				       payload);
			} else {
				printf("progress reacted with %d\n", rc);
				rc = 0;
			}
		} else {
			rc = 0;
		}
		if (payload != NULL)
			free(payload);
		if (topic != NULL)
			free(topic);
		payload = NULL;
		topic = NULL;
	}
exit:

		if (payload != NULL)
			free(payload);
		if (topic != NULL)
			free(topic);
		return rc;
}

static int check_scan_progress(struct EulabeiaScanProgress *progress, char *id){
	int rc;
	char *payload, *topic;
	int payload_len, topic_len;
	while (!eulabeia_scan_finished(progress)){
		if ((rc = mqtt_retrieve_message(
			 &topic, &topic_len, &payload, &payload_len)) == -1) {
			printf("unable to retrieve message, quitting\n");
			rc = -1;
		}
		if (rc == 0) {
			if ((rc = eulabeia_scan_progress(
				 payload, id, progress)) == 0) {
				printf("[%s][scan_id:%s][status:%d] %s\n",
						__FUNCTION__,
				       id,
				       progress->status,
				       payload);
			} else {
				rc = 0;
			}
		} else {
			rc = 0;
		}
		if (payload != NULL)
			free(payload);
		if (topic != NULL)
			free(topic);
		payload = NULL;
		topic = NULL;
	}
exit:

		if (payload != NULL)
			free(payload);
		if (topic != NULL)
			free(topic);
		return rc;
}
int main()
{
	struct EulabeiaScanProgress scanp;
	struct EulabeiaCRUDProgress target_progress, scan_progress;
	struct EulabeiaScan *scan;
	struct EulabeiaTarget *target;
	int rc;

	signal(SIGINT, signalhandler);
	if ((ec = eulabeia_initialize("localhost:9138", NULL)) == NULL) {
		printf("init returned NULL, quitting\n");
		rc = -1;
		goto exit;
	}

	target = example_target();
	printf("creating target %s\n", target->id);
	if ((rc = eulabeia_modify_target(ec, target, GROUP_ID)) != 0){
		printf("[%d] failed to pbulish target\n", rc);
		goto exit;
	}
	target_progress.status = EULABEIA_CRUD_REQUESTED;
	if ((rc = check_for_modify_progress(&target_progress, target->id)) != 0){
		printf("failed (%d) to verify modify target\n", rc);
		goto exit;
	}

	scan = example_scan();
	printf("successfully created target; creating scan %s\n", scan->id);
	if ((rc = eulabeia_modify_scan(ec, scan, GROUP_ID)) != 0){
		printf("[%d] failed to pbulish scan\n", rc);
		goto exit;
	}
	scan_progress.status = EULABEIA_CRUD_REQUESTED;
	if ((rc = check_for_modify_progress(&scan_progress, scan->id)) != 0){
		printf("failed (%d) to verify modify scan\n", rc);
		goto exit;
	}

	printf("successfully created scan; starting scan %s\n", scan->id);

	if ((rc = eulabeia_start_scan(ec, scan, NULL)) != 0) {
		printf("[%d] unable to start scan %s\n", rc, scan->id);
		goto exit;
	}
	if ((rc = check_scan_progress(&scanp, scan->id)) != 0){
		printf("failed (%d) to verify start scan\n", rc);
		goto exit;
	}
exit:
	free_example_target(target);
	free_example_scan(scan);

	eulabeia_destroy(ec);
	return rc;
}
