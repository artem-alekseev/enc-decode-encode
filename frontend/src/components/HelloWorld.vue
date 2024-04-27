<script setup>
import {reactive} from 'vue'
import {OpenEnc, DecompressEnc, CompressEnc} from '../../wailsjs/go/main/App'

const data = reactive({
  resultText: "",
  file: "None",
  xor: {
    first : "92",
    second : "65",
    third : "67",
    four : "57",
  },
})

function openFile() {
  OpenEnc().then(result => {
    data.file = result
  })
}

function compressFile() {
  CompressEnc(data.xor)
}

function decompressFile() {
  DecompressEnc(data.xor)
}

</script>

<template>
  <main style="--wails-draggable:no-drag">
    File
    <button @click="openFile()">Open ENC</button>
    <div id="result" class="result">File: {{ data.file }}</div><br>
    XOR 1 <input v-model="data.xor.first"><br>
    XOR 2 <input v-model="data.xor.second"><br>
    XOR 3 <input v-model="data.xor.third"><br>
    XOR 4 <input v-model="data.xor.four"><br>
    <button @click="compressFile()">Compress</button>
    <button @click="decompressFile()">Decompress</button>
  </main>
</template>

<style scoped>
.result {
  height: 20px;
  line-height: 20px;
  margin: 1.5rem auto;
}

.input-box .btn {
  width: 60px;
  height: 30px;
  line-height: 30px;
  border-radius: 3px;
  border: none;
  margin: 0 0 0 20px;
  padding: 0 8px;
  cursor: pointer;
}

.input-box .btn:hover {
  background-image: linear-gradient(to top, #cfd9df 0%, #e2ebf0 100%);
  color: #333333;
}

.input-box .input {
  border: none;
  border-radius: 3px;
  outline: none;
  height: 30px;
  line-height: 30px;
  padding: 0 10px;
  background-color: rgba(240, 240, 240, 1);
  -webkit-font-smoothing: antialiased;
}

.input-box .input:hover {
  border: none;
  background-color: rgba(255, 255, 255, 1);
}

.input-box .input:focus {
  border: none;
  background-color: rgba(255, 255, 255, 1);
}
</style>
