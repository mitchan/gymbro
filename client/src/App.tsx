import { createSignal, onMount, Show } from "solid-js";
import "./App.css";
import type { JSX } from "solid-js/jsx-runtime";
import { A } from "@solidjs/router";

function App(props: { children?: JSX.Element }) {
  const [loading, setLoading] = createSignal(true);
  const [serverOk, setServerOk] = createSignal(false);

  onMount(() => {
    fetch("http://localhost:8080/api/test")
      .then((resp) => {
        if (resp.ok) {
          return resp.json();
        }

        throw new Error("Generic error");
      })
      .then(() => {
        setServerOk(true);
      })
      .finally(() => {
        setLoading(false);
      });
  });

  return (
    <Show when={!loading() || !serverOk()} fallback="Loading...">
      <div class="p-4 text-center">
        <A href="/">GymBro</A>
      </div>
      {props.children}
    </Show>
  );
}

export default App;
