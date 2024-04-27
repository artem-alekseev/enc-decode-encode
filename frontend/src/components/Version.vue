<script setup>
import {reactive} from 'vue'

const data = reactive({
  currentVersion: "",
  serverVersion: "",
})
</script>

<template>
  <main>
    <b-form>
      <b-form-group>
        <b-form-input type="file"></b-form-input>
      </b-form-group>
    </b-form>

  </main>
</template>

<style scoped>

</style>