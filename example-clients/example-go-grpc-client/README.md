# Example Go client 

## Application summary

Folder can be used as a package and moved into your project as a whole

Grpc files are built and ready for work.

Package exports two functions:

- CreateLogEntry

- GetLogEntriesList

Example usage can be found in the "example" function

### CreateLogEntry

Accepts following parameters:

- UserId (int64)
- Severity (int32)
- LogType (string)
- Section (string)
- Description (string)
- AdditionalData (JSON)
- HappenedAt (Date)

Returns result (type: string), error

### GetLogEntriesList

Accepts following parameters:

- Page (int32) 
- Limit (int32)
- Query (array of pointers to QueryItem from pb_src package)

QueryItem has a form of

<pre>
{
    Param: string
    Value: string
}
</pre>

Param must be one of the following:

- User_id
- Severity
- Log_type
- Section
- Description
- Before
- After

"Before" and "After" are applied to the column Happened_at

Returns entries (type: []*LogEntry), count (type: int64), error

## In case the proto file was updated

Run the following command from the src folder

<pre>
protoc -I api/proto/pb_src api/proto/pb_src/logger-service.proto --go_out=plugins=grpc:api/proto/pb_src
</pre>
