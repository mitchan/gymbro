import { createSignal } from "solid-js";
import { InputText } from "../components/inputs/InputText";
import { Button } from "../components/ui/Button";
import z from "zod";
import { setUser } from "../store/app";
import { useNavigate } from "@solidjs/router";

export default function Register() {
  const navigate = useNavigate();

  const [username, setUsername] = createSignal("");
  const [email, setEmail] = createSignal("");
  const [password, setPassword] = createSignal("");
  const [loading, setLoading] = createSignal(false);

  function onSubmit(e: SubmitEvent) {
    e.preventDefault();

    const userSchema = z.object({
      email: z.email(),
      password: z.string().min(8),
      username: z.string().trim().min(1),
    });

    const user = userSchema.safeParse({
      email: email(),
      password: password(),
      username: username(),
    });

    if (!user.success) {
      // TODO: show errors to the user
      console.error(user.error);
      return;
    }

    setLoading(true);
    fetch("/api/user", {
      method: "POST",
      body: JSON.stringify(user.data),
    })
      .then((resp) => {
        if (!resp.ok) {
          throw new Error("HTTP error");
        }

        return resp.json();
      })
      .then((user) => {
        setUser(user);
        navigate("/", { replace: true });
      })
      .catch((error) => {
        // TODO: show error to the user
        console.error(error);
      })
      .finally(() => {
        setLoading(false);
      });
  }

  return (
    <form class="p-4" on:submit={onSubmit}>
      <InputText
        label="Username"
        name="username"
        value={username()}
        onChange={setUsername}
      />

      <InputText
        label="Email"
        name="email"
        value={email()}
        onChange={setEmail}
      />

      <InputText
        label="Password"
        name="password"
        type="password"
        value={password()}
        onChange={setPassword}
      />

      <Button label="Registrati" type="submit" disabled={loading()} />
    </form>
  );
}
