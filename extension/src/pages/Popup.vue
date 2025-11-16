<script lang="ts" setup>
import { Idle } from "../components/Idle";
import { Footer } from "../components/Footer";
import { Connected } from "../components/Connected";
import { useForm } from "./useForm";

const { connectionState, onConnect, onDisconnect } = useForm();
</script>

<template>
  <div class="size-full">
    <h1
      class="text-center m-4 text-2xl font-semibold tracking-tight text-balance text-white"
    >
      Leetcode reach presence
    </h1>
    <Idle v-if="connectionState !== 'connected'" @connected="onConnect" />
    <Connected v-else @disconnect="onDisconnect" />
    <div class="absolute bottom-4 left-0 w-full flex justify-center">
      <Footer :status="connectionState" />
    </div>
  </div>
</template>
