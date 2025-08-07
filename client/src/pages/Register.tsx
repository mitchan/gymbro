import { createSignal } from "solid-js";
import { InputText } from "../components/inputs/InputText";
import { Button } from "../components/ui/Button";

export default function Register() {
  const [name, setName] = createSignal("");
  const [surname, setSurname] = createSignal("");
  const [email, setEmail] = createSignal("");
  const [password, setPassword] = createSignal("");

  function onSubmit(e: SubmitEvent) {
    e.preventDefault();
  }

  return (
    <form class="p-4" on:submit={onSubmit}>
      <InputText label="Nome" name="name" value={name()} onChange={setName} />

      <InputText
        label="Cognome"
        name="surname"
        value={surname()}
        onChange={setSurname}
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
