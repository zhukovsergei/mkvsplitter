<script setup>
import { reactive } from 'vue'
import { SelectMKVFile, SplitMKV } from '../../wailsjs/go/main/App'

const data = reactive({
  infoText: "Select file below 👇",
  pathToFile: "",
  startTime: "00:00:10",
  endTime: "00:00:20"
})

function selectFile() {
  SelectMKVFile().then(result => {
    data.infoText = result;
    data.pathToFile = result;
  })
}

async function cutFile() {

  if (!data.pathToFile) {
    data.infoText = "Select file!"
    return;
  }

  try {
    const output = await SplitMKV(data.pathToFile, data.startTime, data.endTime)
    data.infoText = "Done: " + output
  } catch (err) {
    data.infoText = "Error: " + err
  }
}
</script>

<template>
  <main>
    <div id="result" class="result">{{ data.infoText }}</div>
    <div id="input" class="input-box">
      <button class="btn" @click="selectFile">Select</button>
    </div>

    <div style="margin-top:20px;">
      <label>
        Start (HH:MM:SS):
        <input type="text" v-model="data.startTime" class="input" style="width:100px;" />
      </label>
      <label style="margin-left:20px;">
        End (HH:MM:SS):
        <input type="text" v-model="data.endTime" class="input" style="width:100px;" />
      </label>

      <button class="btn" @click="cutFile" style="margin-left:20px;">Go</button>
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
  width: 100px;
  height: 30px;
  line-height: 30px;
  border-radius: 3px;
  border: none;
  margin-left: 20px;
  padding: 0 8px;
  cursor: pointer;
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
  background-color: rgba(255, 255, 255, 1);
}

.input-box .input:focus {
  background-color: rgba(255, 255, 255, 1);
}
</style>
