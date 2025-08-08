import { Show } from "solid-js";
import { user } from "../store/app";

export default function Home() {
  return (
    <Show when={user()} fallback={<div>No user</div>}>
      {(user) => {
        return <div>{user().username}</div>;
      }}
    </Show>
  );
}
