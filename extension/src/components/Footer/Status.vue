<script lang="ts" setup>
import { computed } from "vue";
import { ConnectionStatus } from "../../utils/types";

const props = defineProps<{
  status: ConnectionStatus;
}>();

const statusColor = computed(() => {
  switch (props.status) {
    case "connected":
      return "bg-green-500";
    case "disconnected":
      return "bg-red-500";
    default:
      return "bg-orange-400";
  }
});
</script>

<template>
  <span class="inline-flex items-center gap-2">
    <span
      :class="['h-3 w-3 rounded-full animate-pulse', statusColor]"
      aria-label="connection status"
    ></span>
    <span class="text-sm text-white capitalize">{{ status }}</span>
  </span>
</template>
