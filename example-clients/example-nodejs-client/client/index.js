const messages     = require('./logger-service_pb');
const services     = require('./logger-service_grpc_pb');
const timestamp_pb = require('google-protobuf/google/protobuf/timestamp_pb');
const grpc         = require('grpc');
const client       = new services.LoggerServiceClient('0.0.0.0:7777', grpc.credentials.createInsecure());

function createEntry(userId, severity, logType, section, description, additionalData, happenedAt) {
  const logEntry = new messages.LogEntry();

  logEntry.setUserId(userId);
  logEntry.setSeverity(severity);
  logEntry.setLogType(logType);
  logEntry.setSection(section);
  logEntry.setDescription(description);
  logEntry.setAdditionalData(JSON.stringify(additionalData));
  logEntry.setHappenedAt(timestamp_pb.Timestamp.fromDate(happenedAt));


  return createEntryAsync(logEntry)
    .then((data) => {
      return data;
    })
    .catch((error) => {
      return error;
    });
}

function getEntries(page, limit, query) {
  const queryItems = [];

  for (let i = 0; i < query.length; ++i) {
    try {
      const queryItem = new messages.QueryItem;
      queryItem.setParam(query[i].param);
      queryItem.setValue(`${query[i].value}`);

      queryItems.push(queryItem);
    } catch (e) {
      console.log(e);
    }
  }

  const entriesRequest = new messages.EntriesRequest();

  entriesRequest.setPage(page);
  entriesRequest.setLimit(limit);
  entriesRequest.setQueryList(queryItems);

  return getEntriesAsync(entriesRequest)
    .then((data) => {
      return data;
    })
    .catch((error) => {
      return error;
    });
}

module.exports = {
  createEntry,
  getEntries
};

function createEntryAsync(logEntry) {
  return new Promise(function (resolve, reject) {
    return client.createEntry(logEntry, (err, response) => {
      if (err) {
        reject(err);
      }
      resolve(messages.LogCreationResponse.toObject(false, response));
    });
  });
}

function getEntriesAsync(entriesRequest) {
  return new Promise(function (resolve, reject) {
    client.getEntries(entriesRequest, (err, response) => {
      if (err) {
        return reject(err);
      }

      const data = messages.EntriesResponse.toObject(false, response).entriesList;

      for (let i = 0; i < data.length; ++i) {
        if (data[i].happenedAt != null && data[i].happenedAt.seconds != null && data[i].happenedAt.nanos != null) {
          data[i].happenedAt = new Date((data[i].happenedAt.seconds * 1000) + (data[i].happenedAt.nanos / 1000000));
        }
      }

      resolve({
        data,
        count: messages.EntriesResponse.toObject(false, response).count
      });
    });
  });
}



async function example() {
  const queryParams = [{
    param: 'section',
    value: 'Main'
  },{
    param: 'before',
    value: `${Math.round(new Date('2019-05-07 02:08:16').getTime()/1000)}`
  },{
    param: 'after',
    value: '0'
  }];

  const data = await getEntries(1, 20, queryParams);

  console.log(data);
}
