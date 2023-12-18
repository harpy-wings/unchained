import { sha } from "./utils/hash.js";
import {
  State,
  Keys,
  Config,
  StringAnyObject,
  StringGossipMethodObject,
} from "./types.js";

export const version = "0.8.0";
export const protocolVersion = "0.8.0";

export const topic = sha(`Kenshi.Unchained.Testnet.Topic.V${protocolVersion}`);

export const sockets = new Map();
export const keys: Keys = {};

export const config: Config = {
  name: "Change Me",
  log: "info",
  rpc: {
    ethereum: "https://ethereum.publicnode.com",
  },
  database: {
    url: "",
    name: "unchained",
  },
  secretKey: "",
  lite: false,
};

export const rpcMethods: StringAnyObject = {};
export const gossipMethods: StringGossipMethodObject<any> = {};

export const errors = {
  E_NOT_FOUND: 404,
  E_INTERNAL: 500,
};

export const state: State = {
  connected: false,
};

export const nameRegex = /^[a-z0-9 @\._'-]+$/i;
