import { createSignal } from "solid-js";
import z from "zod";

// TODO: Move to schema file?
export const userSchema = z.object({
  id: z.uuidv4(),
  username: z.string(),
});

type User = z.infer<typeof userSchema>;

export const [user, setUser] = createSignal<User | null>(null);
