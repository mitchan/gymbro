import { createSignal } from "solid-js";
import { InputText } from "../components/inputs/InputText";
import { Button } from "../components/ui/Button";
import z from "zod";
import { setUser } from "../store/app";
import { useNavigate } from "@solidjs/router";

export default function Login() {
  const navigate = useNavigate();

  const [email, setEmail] = createSignal("");
  const [password, setPassword] = createSignal("");
  const [loading, setLoading] = createSignal(false);

  const loginSchema = z.object({
    email: z.email(),
    password: z.string().min(8),
  });

  function onSubmit(e: SubmitEvent) {
    e.preventDefault();

    const loginData = loginSchema.safeParse({
      email: email(),
      password: password(),
    });

    if (!loginData.success) {
      // TODO: show errors to the user
      console.error(loginData.error);
      return;
    }

    setLoading(true);
    fetch("/api/user/login", {
      method: "POST",
      body: JSON.stringify(loginData.data),
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
