<script setup>
import {reactive} from 'vue'
import {CloseApp, CompressEnc, DecompressEnc, OpenEnc} from '../../wailsjs/go/main/App'

const data = reactive({
  resultText: "",
  file: "Open file",
  xor: {
    first: "92",
    second: "65",
    third: "67",
    four: "57",
  },
})

function closeApp() {
  CloseApp()
}

function openFile() {
  OpenEnc().then(result => {
    data.file = result
  })
}

function compressFile() {
  CompressEnc(data.xor).then(result => {
    data.file = result
  })
}

function decompressFile() {
  DecompressEnc(data.xor).then(result => {
    data.file = result
  })
}

</script>

<template>
  <main style="--wails-draggable:no-drag" class="container">
    <div id="result" class="result">{{ data.file }}</div>
    <div class=" btn-group btn-group-lg">
      <button class="btn btn-success" @click="openFile()">Open ENC \ DEC</button>
      <button class="btn btn-success" @click="compressFile()">Compress</button>
      <button class="btn btn-success" @click="decompressFile()">Decompress</button>
      <button class="btn btn-danger" @click="closeApp()">Close</button>
    </div>
    <div class="form-group mt-2">
      <label for="xor1">XOR 1</label>
      <input id="xor1" type="text" class="form-control small" v-model="data.xor.first">
      <label for="xor2">XOR 2</label>
      <input id="xor2" type="text" class="form-control" v-model="data.xor.second">
      <label for="xor3">XOR 3</label>
      <input id="xor3" type="text" class="form-control" v-model="data.xor.third">
      <label for="xor4">XOR 4</label>
      <input id="xor4" type="text" class="form-control" v-model="data.xor.four">
    </div>
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
