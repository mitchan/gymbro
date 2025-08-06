/* @refresh reload */
import { render } from "solid-js/web";
import { Route, Router } from "@solidjs/router";
import "./index.css";
import App from "./App";
import { lazy } from "solid-js";

const root = document.getElementById("root");

const Register = lazy(() => import("./pages/Register"));

if (root) {
  render(
    () => (
      <Router root={App}>
        <Route path="/" />
        <Route path="/register" component={Register} />
      </Router>
    ),
    root
  );
}
