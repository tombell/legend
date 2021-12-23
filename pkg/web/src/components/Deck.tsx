import { h } from "preact";

interface Props {
  id: number;
  artist: string;
  name: string;
}

export default ({
  id,
  artist,
  name,
}: Props) => (
  <div key={`deck-${id}`} class="text-white">
    <div class="container ml-0 mt-3">
      <div class="col-3">
        <h6
          class="font-weight-bold text-center w-50"
          style={{ background: "#ff0072" }}
        >
          {`DECK ${["ONE", "TWO"][id]}`}
        </h6>
      </div>
    </div>

    <div class="container ml-0">
      <div class="col-12">
        <h4 class="font-weight-bold">{artist}</h4>
        <h5 class="mt-n1">{name}</h5>
      </div>
    </div>
  </div>
);
