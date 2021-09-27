if(description)
{
  script_oid("0.0.0.0.0.0.0.0.0.1");
  script_name("keys");
  script_family("my test family");  

  exit(0);
}

set_kb_item( name: "test/key1", value: TRUE );
set_kb_item( name: "test/key2", value: 42 );
set_kb_item( name: "test/key3", value: "waldfee" );
