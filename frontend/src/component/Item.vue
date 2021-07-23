<template>
  <div class="shadow item">
    <!-- Icon -->
    <div style="width: 80px; display: flex; align-items: center; justify-content: center">
      <img
        v-if="item.type === 'credential'"
        src="../asset/lock.svg"
        alt="Settings"
        draggable="false"
      />
      <img v-if="item.type === 'token'" src="../asset/key.svg" alt="Settings" draggable="false" />
      <img v-if="item.type === 'text'" src="../asset/text.svg" alt="Settings" draggable="false" />
    </div>

    <div style="flex: 1; margin-left: auto; display: flex; flex-direction: column">
      <div class="item-header">
        <div class="item-service" style="margin-right: auto">{{ item.service }}</div>
        <img
          @click="$emit('edit', item)"
          class="clickable"
          src="../asset/pencil.svg"
          alt="Pencil"
          draggable="false"
        />
        <img
          @click="copyToBuffer(item.content[0])"
          class="clickable"
          src="../asset/copy.svg"
          alt="Copy"
          draggable="false"
        />
        <img
          @click="copyToBuffer(item.content[1])"
          v-if="item.type === 'credential'"
          class="clickable"
          src="../asset/copy.svg"
          alt="Copy"
          draggable="false"
        />
      </div>
      <div class="item-description">{{ item.description }}</div>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';

export default defineComponent({
  props: {
    item: Object,
  },
  components: {},
  async mounted() {},
  methods: {
    copyToBuffer(data: string) {
      var tempInput = document.createElement('input');
      tempInput.value = data;
      document.body.appendChild(tempInput);
      tempInput.select();
      document.execCommand('copy');
      document.body.removeChild(tempInput);
    },
  },
  data: () => {
    return {};
  },
});
</script>

<style lang="scss" scoped>
.item {
  background: #3d3d3d;
  display: flex;
  height: 100px;
  border-radius: 6px;
  overflow: hidden;
  flex: none;

  img {
    user-select: none;
  }

  .item-header {
    background: #515151;
    padding: 10px 15px;
    color: #a4a4a4;
    font-size: 16px;
    font-weight: bold;
    text-transform: uppercase;
    display: flex;
    align-items: center;
    border-bottom: 1px solid #313131;
    border-left: 1px solid #313131;

    .item-service {
      max-width: 180px;
      overflow: hidden;
      text-overflow: ellipsis;
    }

    > img {
      margin-left: 10px;
    }
  }

  .item-description {
    overflow: hidden;
    text-overflow: ellipsis;
    word-break: break-all;
  }

  .item-description {
    flex: 1;
    background: #464646;
    color: #8b8b8b;
    padding: 10px 15px;
    border-left: 1px solid #313131;
  }
}
</style>
