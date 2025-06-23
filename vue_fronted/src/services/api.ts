// api.js - 创建axios实例并配置baseURL
import axios from 'axios';

// 创建API实例
const apiClient = axios.create({
  baseURL: 'http://localhost:8082', // 配置基础URL
  timeout: 10000,
  headers: {
    'Content-Type': 'application/json',
  },
});

// 请求拦截器（可选）
apiClient.interceptors.request.use(
  (config) => {
    // 可在此添加认证信息
    return config;
  },
  (error) => {
    return Promise.reject(error);
  }
);

// 响应拦截器（可选）
apiClient.interceptors.response.use(
  (response) => response.data,
  (error) => {
    // 统一处理错误
    console.error('API请求错误:', error);
    return Promise.reject(error?.response?.data || error.message);
  }
);
// 导出API实例
export default apiClient;