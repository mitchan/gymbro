import { createSignal } from "solid-js";
import { InputText } from "../components/inputs/InputText";
import { Button } from "../components/ui/Button";

export default function Register() {
  const [username, setUsername] = createSignal("");
  const [email, setEmail] = createSignal("");
  const [password, setPassword] = createSignal("");

  function onSubmit(e: SubmitEvent) {
    e.preventDefault();

    fetch("http://localhost:8080/api/user", {
      method: "POST",
      body: JSON.stringify({
        email: email(),
        password: password(),
        username: username(),
      }),
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
        value={password()}
        onChange={setPassword}
      />

      <Button label="Registrati" type="submit" />
    </form>
  );
}
