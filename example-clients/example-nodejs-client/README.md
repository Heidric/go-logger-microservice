# Example Node.js client 

## Application summary

All of the required packages are listed in the package.json file.

Files are built and are ready for work.

Folder "client" exports two functions:

- createEntry

- getEntries

Example getEntries usage can be found in the "example" function

### createEntry

Accepts following parameters:

- userId (integer)
- severity (integer)
- logType (string)
- section (string)
- description (string)
- additionalData (JSON)
- happenedAt (Date)

### getEntries

Accepts following parameters:

- page (integer) 
- limit (integer)
- query (array of objects containing fields "param" and "value")

Param must be one of the following:

- user_id
- severity
- log_type
- section
- description
- before
- after

"before" and "after" are applied to the column happened_at

Returns an object of the following form:

<pre>
{
  data:
   [
     { userId: 1,
       severity: 1,
       logType: 'New',
       section: 'Main',
       description: 'Qwerty',
       additionalData: '{"info":{"a":"b"}}',
       happenedAt: 2019-05-06T17:16:19.624Z
     }
   ],
  count: 1
}
</pre>

## In case the proto file was updated

Execute the following bash script

```sh
npm i -g grpc-tools
grpc_tools_node_protoc --js_out=import_style=commonjs,binary:./client/ --grpc_out=./client --plugin=protoc-gen-grpc=`which grpc_tools_node_protoc_plugin` logger-service.proto
```
