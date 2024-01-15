import { sockets, rpcMethods } from "../constants.js";
import { MetaData } from "../types.js";
import { brotliCompressSync } from "zlib";

export interface WantPacket {
  dataset: string;
  want: string;
  have: string[];
}

export interface WantAnswer {
  dataset: string;
  want: string;
  have: any[];
}

export interface Dataset {
  want: (data: WantPacket) => any[] | Promise<any[]>;
  have: (data: WantAnswer) => void | Promise<void>;
}

export const datasets = new Map<string, Dataset>();

const writePayload = (nodes: MetaData[], payload: Buffer) => {
  for (const node of nodes) {
    if (!node.socket.closed) {
      const sent = node.socket.write(payload);
      // TODO: Maybe add back pressure config to allow N cached messages
      if (!sent) {
        node.needsDrain = true;
      }
    }
  }
};

const wantRpcCall = (nodes: MetaData[], data: WantPacket) => {
  const call = { type: "call", request: { method: "want", args: data } };
  const payload = brotliCompressSync(JSON.stringify(call));
  writePayload(nodes, payload);
};

const haveRpcCall = async (nodes: MetaData[], data: WantAnswer) => {
  const call = { type: "call", request: { method: "have", args: data } };
  const payload = brotliCompressSync(JSON.stringify(call));
  writePayload(nodes, payload);
};

export const queryNetworkFor = (
  want: string,
  dataset: string,
  have: string[] = []
) => {
  const nodes = [...sockets.values()];
  const packet: WantPacket = { want, dataset, have };
  wantRpcCall(nodes, packet);
};

const want = async (data: WantPacket, sender: MetaData) => {
  const dataset = datasets.get(data.dataset);
  if (!dataset) {
    return;
  }
  const have = await dataset.want(data);
  if (!have.length) {
    return;
  }
  const packet: WantAnswer = { dataset: data.dataset, want: data.want, have };
  haveRpcCall([sender], packet);
};

const have = async (data: WantAnswer) => {
  const dataset = datasets.get(data.dataset);
  if (!dataset) {
    return;
  }
  await dataset.have(data);
};

Object.assign(rpcMethods, { want, have });