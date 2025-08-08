import { createSignal } from "solid-js";

interface User {
  AccessToken: string;
  id: string;
  username: string;
}

export const [user, setUser] = createSignal<User | null>(null);
