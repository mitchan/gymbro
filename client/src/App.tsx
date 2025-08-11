import { A } from "@solidjs/router";
import "./App.css";
import type { JSX } from "solid-js/jsx-runtime";

function App(props: { children?: JSX.Element }) {
  return (
    <>
      <div class="p-4 text-center">
        <A href="/">GymBro</A>
      </div>
      {props.children}
    </>
  );
}

export default App;
