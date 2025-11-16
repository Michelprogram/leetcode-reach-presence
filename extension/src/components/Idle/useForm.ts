import { ref } from "vue";

type Options = {
  evt: (e: "connected", port: string) => void;
};

export const useForm = (options: Options) => {
  const port = ref<string>("8085");
  const error = ref<string>("");

  const onOpen = async (event: Event) => {
    if (error.value !== "") return;

    const socket = event.target as WebSocket;

    options.evt("connected", port.value);

    socket.close();
  };

  const onError = () => {
    error.value = "Invalid port";
    console.warn("Error connecting to server on port:", port.value);
  };

  const onConnection = () => {
    error.value = "";
    const socket = new WebSocket(`ws://localhost:${port.value}`);
    socket.addEventListener("error", onError);
    socket.addEventListener("open", onOpen);
  };

  return { port, onConnection, error };
};
