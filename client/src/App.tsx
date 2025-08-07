import { createSignal, onMount, Show } from "solid-js";
import "./App.css";
import type { JSX } from "solid-js/jsx-runtime";
import { A } from "@solidjs/router";

import * as rd from "@devexperts/remote-data-ts";

function App(props: { children?: JSX.Element }) {
  const [serverOk, setServerOk] = createSignal<rd.RemoteData<unknown, boolean>>(
    rd.initial
  );

  onMount(() => {
    setServerOk(rd.pending);
    fetch("http://localhost:8080/api/test")
      .then((resp) => {
        if (resp.ok) {
          return resp.json();
        }

        throw new Error("Generic error");
      })
      .then(() => {
        setServerOk(rd.success(false));
      })
      .catch((error) => {
        rd.failure(error);
      });
  });

  return (
    <Show when={rd.isSuccess(serverOk())} fallback="Loading...">
      <div class="p-4 text-center">
        <A href="/">GymBro</A>
      </div>
      {props.children}
    </Show>
  );
}

export default App;
