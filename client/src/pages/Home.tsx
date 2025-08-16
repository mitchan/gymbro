import { createSignal, For, onMount, Show } from "solid-js";
import { apiClient } from "../lib/api/apiClient";

export default function Home() {
  const [workouts, setWorkouts] = createSignal<unknown[]>([]);

  onMount(() => {
    apiClient
      .fetch("/api/workouts")
      .then((workouts) => {
        if (Array.isArray(workouts)) {
          // TODO: validate workouts
          setWorkouts(workouts);
        }
      })
      .catch((error) => {
        // TODO: handle error
        console.error(error);
        setWorkouts([]);
      });
  });

  return (
    <Show when={workouts().length > 0} fallback={<div>No workouts</div>}>
      <ul>
        <For each={workouts()}>{() => <div>todo</div>}</For>
      </ul>
    </Show>
  );
}
