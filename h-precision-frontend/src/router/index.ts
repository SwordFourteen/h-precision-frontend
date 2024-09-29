import { createRouter, createWebHistory } from 'vue-router';

import Login from '../views/Login.vue';
import Home from '../views/Home.vue';
import ToolStrategyManagement from '../views/ToolStrategyManagement.vue';
import ToolScriptManagement from '../views/ToolScriptManagement.vue';
import ModelTuning from '../views/ModelTuning.vue';
import VulnerabilityDetection from '../views/VulnerabilityDetection.vue';

const routes = [
  { path: '/', component: Login },
  { path: '/home', component: Home },
  { path: '/tool-strategy', component: ToolStrategyManagement },
  { path: '/tool-script', component: ToolScriptManagement },
  { path: '/model-tuning', component: ModelTuning },
  { path: '/vulnerability-detection', component: VulnerabilityDetection },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

export default router;
