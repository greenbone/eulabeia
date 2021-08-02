Error Handling
==============

In general `eulabeia` sends a [info.Failure](/messages/info/info.go):

```
{
  "id": "f704d1e0-768d-4a86-ab6a-dd28e3f45776",
  "created": 443397956,
  "message_type": "failure.modify.target",
  "message_id": "363cde52-e07f-11eb-99c4-6b7f958f017",
  "group_id": "12",
  "error": "Insufficient space",
}
```

when an error occured while retrieving / already executing a cmd.

error happening without wrongful usage
--------------------------------------

This error occurs even when the usage is correct.

For this case most of the modules are written to return the error so that the cmd can handle the case explicitely.

For an instance if a director cannot store a target because there is insufficient space on the device it should:

1.	send a [info.Failure](/messages/info/info.go) response so that the client knows that storing failed
2.	log the error message
3.	exit director with an error code so that it is escalated to the infrastructure (e.g. restart systemd, container)

response on wrongful usage
--------------------------

This error occurs when a cmd event is containing either:

-	wrong values on cmd.Modify
-	wrong ID on cmd.Start, cmd.Get, cmd.stop, ...

in this case `eulabeia` will not return an error to the module user but rather

1.	send a [info.Failure](/messages/info/info.go) response to that the client can resend a correct cmd

`eulabeia` does not handle unknown messages as error but is returning a [info.Failure](/messages/info/info.go) when the topic is correct and the message contains the eulabeia meta-data; otherwise this message will be ignored.

This allows to run different version of `eulabeia` in parallel without having to concern about new cmds and it allows deployment methods like blue/green on a wide arrange sensors wihtout having to deal with restart spam but still having the chance to monitor weird behaviour when e.g. a sensor wasn't updated but should be.
