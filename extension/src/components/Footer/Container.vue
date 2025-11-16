<script setup lang="ts">
import Status from "./Status.vue";
import { ConnectionStatus } from "../../utils/types";

const props = defineProps<{
  status: ConnectionStatus;
}>();
</script>

<template>
  <footer
    class="w-full flex justify-between items-center p-4 text-white text-sm"
  >
    <div>
      <Status :status="props.status" />
    </div>
    <div>
      <span>
        <i>
          <a
            href="https://github.com/michelprogram"
            target="_blank"
            rel="noopener"
            class="hover:underline"
            >michelprogram</a
          >
        </i>
      </span>
    </div>
  </footer>
</template>
