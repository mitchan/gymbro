import { createSignal, onMount, Show } from "solid-js";
import "./App.css";

function App() {
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
        console.log("asdas");
        setServerOk(true);
      })
      .finally(() => {
        setLoading(false);
      });
  });

  return (
    <Show when={!loading()} fallback="Loading...">
      <div>Server is: {serverOk() ? "Online" : "Offline"}</div>
    </Show>
  );
}

export default App;
