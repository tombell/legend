/* eslint-disable no-console */

import { useEffect } from 'preact/hooks';

export default function useWebSocket(
  host: string,
  onmessage: (ev: MessageEvent) => void
) {
  useEffect(() => {
    const socket = new WebSocket(host);

    socket.onopen = (e) => {
      console.log('connected to websocket server', e);
    };

    socket.onclose = (e) => {
      if (e.wasClean) {
        console.log('disconnected from websocket server', e);
      } else {
        console.log('connection to websocket server died', e);
      }
    };

    socket.onmessage = onmessage;
  }, []);
}
