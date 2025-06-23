const FRONT_TOKEN_KEY = 'front_token';
const ADMIN_TOKEN_KEY = 'admin_token';
const FRONT_USERNAME_KEY = 'front_username';
const ADMIN_USERNAME_KEY = 'admin_username';

class Auth {
  // 添加类型注解
  frontToken: string | null;
  frontUsername: string | null;
  adminToken: string | null;
  adminUsername: string | null;

  constructor() {
    this.frontToken = localStorage.getItem(FRONT_TOKEN_KEY);
    this.frontUsername = localStorage.getItem(FRONT_USERNAME_KEY);
    this.adminToken = localStorage.getItem(ADMIN_TOKEN_KEY);
    this.adminUsername = localStorage.getItem(ADMIN_USERNAME_KEY);
  }

  // 添加类型注解
  setFrontAuth(token: string, username: string): void {
    localStorage.setItem(FRONT_TOKEN_KEY, token);
    localStorage.setItem(FRONT_USERNAME_KEY, username);
    this.frontToken = token;
    this.frontUsername = username;
  }

  // 添加类型注解
  setAdminAuth(token: string, username: string): void {
    localStorage.setItem(ADMIN_TOKEN_KEY, token);
    localStorage.setItem(ADMIN_USERNAME_KEY, username);
    this.adminToken = token;
    this.adminUsername = username;
  }

  // 修复方法名拼写错误
  deleteFrontAuth(): void {
    localStorage.removeItem(FRONT_TOKEN_KEY);
    localStorage.removeItem(FRONT_USERNAME_KEY);
    this.frontToken = null;
    this.frontUsername = null;
  }

  // 修复方法名拼写错误
  deleteAdminAuth(): void {
    localStorage.removeItem(ADMIN_TOKEN_KEY);
    localStorage.removeItem(ADMIN_USERNAME_KEY);
    this.adminToken = null;
    this.adminUsername = null;
  }

  // 添加常用方法：检查前端用户是否已认证
  isFrontAuthenticated(): boolean {
    return !!this.frontToken;
  }

  // 添加常用方法：检查管理员是否已认证
  isAdminAuthenticated(): boolean {
    return !!this.adminToken;
  }
}

const auth = new Auth();
export default auth;