import { onMount, Show } from "solid-js";
import type { JSX } from "solid-js/jsx-runtime";
import { setUser, user, userSchema } from "../store/app";

interface Props {
  children?: JSX.Element;
}

export default function UserProtected(props: Props) {
  // TODO: show loading
  onMount(() => {
    if (user()) {
      return;
    }

    fetch("/api/user/me", {
      method: "get",
      credentials: "include",
    })
      .then((resp) => {
        if (!resp.ok) {
          throw new Error("HTTP error");
        }

        return resp.json();
      })
      .then((user) => {
        const userData = userSchema.safeParse(user);
        if (!userData.success) {
          throw new Error("invalid user data");
        }
        setUser(userData.data);
      })
      .catch((error) => {
        // TODO: show error to the user
        console.error(error);
      })
      .finally(() => {});
  });

  return <Show when={user()}>{props.children}</Show>;
}
