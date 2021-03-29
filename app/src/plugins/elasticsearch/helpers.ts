import { formatTime } from 'utils/helpers';

// ITimes is the interface for a start and end time.
export interface ITimes {
  timeEnd: number;
  timeStart: number;
}

// IElasticsearchOptions is the interface for all options, which can be set for an Elasticsearch query.
export interface IElasticsearchOptions extends ITimes {
  fields?: string[];
  query: string;
  queryName: string;
  scrollID?: string;
}

// IDocument is the interface for a single Elasticsearch document. It is just an general interface for the JSON
// representation of this document.
export interface IDocument {
  // eslint-disable-next-line @typescript-eslint/no-explicit-any
  [key: string]: any;
}

// getOptionsFromSearch is used to get the Elasticsearch options from a given search location.
export const getOptionsFromSearch = (search: string): IElasticsearchOptions => {
  const params = new URLSearchParams(search);
  const fields = params.getAll('field');
  const query = params.get('query');
  const timeEnd = params.get('timeEnd');
  const timeStart = params.get('timeStart');

  return {
    fields: fields.length > 0 ? fields : undefined,
    query: query ? query : '',
    queryName: '',
    scrollID: '',
    timeEnd: timeEnd ? parseInt(timeEnd as string) : Math.floor(Date.now() / 1000),
    timeStart: timeStart ? parseInt(timeStart as string) : Math.floor(Date.now() / 1000) - 3600,
  };
};

// getFieldsRecursively returns the fields for a single document as a list of string.
export const getFieldsRecursively = (prefix: string, document: IDocument): string[] => {
  const fields: string[] = [];
  for (const field in document) {
    if (typeof document[field] === 'object') {
      fields.push(...getFieldsRecursively(prefix ? `${prefix}.${field}` : field, document[field]));
    } else {
      fields.push(prefix ? `${prefix}.${field}` : field);
    }
  }

  return fields;
};

// getFields is used to get all fields as strings for the given documents. To get the fields we are looping over the
// given documents and adding each field from this document. As a last step we have to remove all duplicated fields.
export const getFields = (documents: IDocument[]): string[] => {
  const fields: string[] = [];
  for (const document of documents) {
    fields.push(...getFieldsRecursively('', document['_source']));
  }

  const uniqueFields: string[] = [];
  for (const field of fields) {
    if (uniqueFields.indexOf(field) === -1) {
      uniqueFields.push(field);
    }
  }

  return uniqueFields;
};

// getProperty returns the property of an object for a given key.
// eslint-disable-next-line @typescript-eslint/no-explicit-any, @typescript-eslint/explicit-module-boundary-types
export const getProperty = (obj: any, key: string): string | number => {
  return key.split('.').reduce((o, x) => {
    return typeof o == 'undefined' || o === null ? o : o[x];
  }, obj);
};

// formatTimeWrapper is a wrapper for our shared formatTime function. It is needed to convert a given time string to the
// corresponding timestamp representation, which we need for the formatTime function.
export const formatTimeWrapper = (time: string): string => {
  return formatTime(Math.floor(new Date(time).getTime() / 1000));
};