export interface IClusters {
  [satellite: string]: ICluster[];
}

export interface ICluster {
  id: string;
  cluster: string;
  satellite: string;
  updatedAt: number;
}

export interface INamespaces {
  [cluster: string]: INamespace[];
}

export interface INamespace {
  id: string;
  namespace: string;
  cluster: string;
  satellite: string;
  clusterID: string;
  updatedAt: number;
}

export interface IResource {
  id: string;
  description: string;
  isCRD: boolean;
  path: string;
  resource: string;
  scope: string;
  title: string;
  columns: ICRDColumn[];
}

export interface ICRDColumn {
  description: string;
  jsonPath: string;
  name: string;
  type: string;
}