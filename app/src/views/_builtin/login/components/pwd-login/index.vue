<template>
  <n-form ref="formRef" :model="model" :rules="rules" size="large" :show-label="false">
    <n-form-item path="userName">
      <n-input v-model:value="model.userName" placeholder="请输入用户名" @keyup.enter="jumpToPwd" />
    </n-form-item>
    <n-form-item path="password">
      <n-input
        v-model:value="model.password"
        type="password"
        show-password-on="click"
        placeholder="请输入密码"
				ref="pwdInstRef"
        @keyup.enter="jumpToCode"
      />
    </n-form-item>
    <n-form-item path="code">
      <div class="flex flex-row">
        <n-input v-model:value="model.code" placeholder="请输入验证码" @keyup.enter="handleSubmit" ref="codeInstRef" />
        <n-image :src="codeImg" preview-disabled @click="fecthCodeImg"></n-image>
      </div>
    </n-form-item>
    <n-space :vertical="true" :size="24">
      <div class="flex-y-center justify-between">
        <n-checkbox v-model:checked="rememberMe">记住我</n-checkbox>
        <n-button :text="true" @click="toLoginModule('reset-pwd')">忘记密码？</n-button>
      </div>
      <n-button
        type="primary"
        size="large"
        :block="true"
        :round="true"
        :loading="auth.loginLoading"
        @click="handleSubmit"
      >
        确定
      </n-button>
      <!-- <div class="flex-y-center justify-between">
        <n-button class="flex-1" :block="true" @click="toLoginModule('code-login')">
          {{ loginModuleLabels['code-login'] }}
        </n-button>
        <div class="w-12px"></div>
        <n-button class="flex-1" :block="true" @click="toLoginModule('register')">
          {{ loginModuleLabels.register }}
        </n-button>
      </div> -->
    </n-space>
    <!-- <other-account @login="handleLoginOtherAccount" /> -->
  </n-form>
</template>

<script setup lang="ts">
import { reactive, ref, onMounted } from 'vue';
import type { FormInst, FormRules } from 'naive-ui';
import dayjs from 'dayjs';
import axios from 'axios';
import { loginModuleLabels } from '@/constants';
import { useAuthStore } from '@/store';
import { useRouterPush } from '@/composables';
import { formRules } from '@/utils';
import { getServiceEnvConfig, imgBaseUrl } from '~/.env-config';
import { OtherAccount } from './components';
import { InputInst } from 'naive-ui'
import { createRequiredFormRule } from '@/utils';

const { url, proxyPattern } = getServiceEnvConfig(import.meta.env);
const isHttpProxy = import.meta.env.VITE_HTTP_PROXY === 'Y';

const baseURL = isHttpProxy ? proxyPattern : url;
const codeImg = ref('');
let codeToken = '';

const auth = useAuthStore();
const { login } = useAuthStore();
const { toLoginModule } = useRouterPush();

const formRef = ref<HTMLElement & FormInst>();

const model = reactive({
  userName: '',
  password: '',
  code: ''
});

const rules: FormRules = {
	userName: createRequiredFormRule('请输入用户名'),
  password: createRequiredFormRule('请输入密码'),
	code: createRequiredFormRule('请输入验证码'),
};

const pwdInstRef = ref<InputInst >();
const codeInstRef = ref<InputInst >();

const rememberMe = ref(false);

async function handleSubmit() {
  await formRef.value?.validate();

  const { userName, password, code } = model;

  login(userName, password, code, codeToken);
}

function handleLoginOtherAccount(param: { userName: string; password: string }) {
  const { userName, password } = param;
  login(userName, password);
}

function fecthCodeImg() {
  axios
    .get(`${baseURL}/user/codeImage?time=${dayjs().unix()}`, {
      responseType: 'arraybuffer',
      headers: {
        'Content-Type': 'image/png'
      }
    })
    .then(resp => {
      codeImg.value = window.URL.createObjectURL(new Blob([resp.data]));
      codeToken = resp.headers.codetoken
    })
    .catch(e => {
      console.log(e);
    });
}

onMounted(() => {
  fecthCodeImg();
});

function jumpToPwd() {
	console.log(pwdInstRef);

	pwdInstRef.value?.focus();
}

function jumpToCode() {
	codeInstRef.value?.focus();
}
</script>

<style scoped></style>
