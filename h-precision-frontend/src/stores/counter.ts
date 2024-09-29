import { defineStore } from 'pinia';

export const useMainStore = defineStore('main', {
  state: () => ({
    user: null,
    tools: [],
    strategies: [],
    // 其他状态
  }),
  actions: {
    // 定义同步或异步的状态修改方法
  },
});
