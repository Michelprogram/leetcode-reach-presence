<script lang="ts" setup>
import { useForm } from "./useForm";

const emit = defineEmits<{
  (e: "connected", port: string): void;
}>();

const { port, error, onConnection } = useForm({ evt: emit });
</script>

<template>
  <div class="flex flex-col gap-4 p-3 items-center">
    <div>
      <label for="port" class="block text-sm/6 font-medium text-white"
        >Port
        <span
          :class="[
            'transition-opacity',
            'duration-300',
            error ? 'opacity-100 text-red-400' : 'opacity-0',
          ]"
          >({{ error }})</span
        ></label
      >
      <div class="mt-2">
        <div
          class="flex items-center rounded-md bg-white/5 pl-3 outline-1 -outline-offset-1 outline-white/10 focus-within:outline-2 focus-within:-outline-offset-2 focus-within:outline-indigo-500"
        >
          <input
            v-model="port"
            id="port"
            type="text"
            name="port"
            placeholder="default 8085"
            class="block min-w-0 grow bg-transparent py-1.5 pr-3 pl-1 text-base text-white placeholder:text-gray-500 focus:outline-none"
          />
        </div>
      </div>
    </div>
    <button
      @click="onConnection"
      class="inline-flex items-center justify-center border align-middle select-none font-sans font-medium text-center transition-all duration-300 ease-in disabled:opacity-50 disabled:shadow-none disabled:cursor-not-allowed data-[shape=pill]:rounded-full data-[width=full]:w-full focus:shadow-none text-sm rounded-md py-2 px-4 shadow-sm hover:shadow-md bg-slate-800 border-slate-800 text-slate-50 hover:bg-slate-700 hover:border-slate-700"
    >
      Connect
    </button>
  </div>
</template>
