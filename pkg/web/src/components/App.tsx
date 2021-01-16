import { h, Fragment } from 'preact';
import { useState } from 'preact/hooks';

import Deck from 'components/Deck';

import useWebSocket from 'hooks/useWebSocket';

interface Track {
  number: number;
  artist: string;
  name: string;
}

export default () => {
  const [tracks, setTracks] = useState<Track[]>([
    // { number: 1, artist: 'Igor Gonya', name: 'Numbness (Original Mix)' },
    // { number: 2, artist: 'Roberto Surace', name: 'Joys (Purple Disco Machine Extended Remix)' },
  ]);

  useWebSocket('ws://192.168.1.103:8888', (e) => {
    const { history }: { history: Track[] } = JSON.parse(e.data);

    if (history.length === 1) {
      setTracks(history);
      return;
    }

    if (history.length > 1) {
      const [a, b] = history.splice(-2);
      setTracks(a.number % 2 === 0 ? [b, a] : [a, b]);
    }
  });

  return (
    <Fragment>
      {tracks.map((track, idx) => (
        <Deck id={idx + 1} artist={track.artist} name={track.name} />
      ))}
    </Fragment>
  );
};
