if(description)
{
  script_oid("1.3.6.1.4.1.25623.1.0.90022");
  script_version("2019-11-10T15:30:28+0000");
  script_name("mqtt test");
  script_category(ACT_INIT);
  script_family("my test family");  
  script_tag(name:"some", value:"tag");
  script_tag(name:"last_modification", value:"2019-11-10 15:30:28 +0000 (Tue, 10 Nov 2020)");
  script_tag(name:"creation_date", value:"2015-03-27 12:00:00 +0100 (Fri, 27 Mar 2015)");
  script_tag(name:"cvss_base", value:"0.0");
  script_tag(name:"cvss_base_vector", value:"AV:N/AC:L/Au:N/C:N/I:N/A:N");
  script_tag(name:"summary", value:"Report functions were used.");
  script_tag(name:"qod_type", value:"remote_app");

  exit(0);
}

sec_msg = "this is a security message";
log_msg = "this is a log message";
err_msg = "this is a error message";

security_message(data:sec_msg);
log_message(data:log_msg);
error_message(data:err_msg);

exit(0);
