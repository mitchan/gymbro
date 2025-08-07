interface Props {
  type?: "text" | "password" | "number" | "hidden";
  label: string;
  name: string;
  value?: string | number;
  disabled?: boolean;
  hideLabel?: boolean;
  onChange?: (value: string) => void;
  extraContainerClasses?: string;
}

export function InputText(props: Props) {
  const { type = "text" } = props;

  return (
    <div class={`mb-1 flex flex-col ${props.extraContainerClasses ?? ""}`}>
      {type !== "hidden" && !props.hideLabel ? (
        <label>{props.label}</label>
      ) : null}
      <input
        type={type}
        name={props.name}
        value={props.value === 0 ? "" : props.value}
        disabled={props.disabled}
        class="rounded border border-solid border-gray-400 p-2 text-black"
        onChange={(e) => {
          props.onChange?.(e.target.value);
        }}
        placeholder={props.label}
      />
    </div>
  );
}
