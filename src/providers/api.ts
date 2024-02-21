import axios, { AxiosRequestConfig } from "axios";
import { errorInterceptor } from "./interceptors/errorInterceptor";

export const BASE_URL = process.env.NEXT_PUBLIC_BASE_URL || "http://localhost:8000";

class ApiProvider {
  private instance;

  constructor() {
    this.instance = axios.create({
      baseURL: BASE_URL,
    });

    this.instance.interceptors.response.use(
      (response) => response,
      (error) => errorInterceptor(error),
    );
  }
  async get<T>(path: string, options?: AxiosRequestConfig): Promise<T> {
    try {
      const response = await this.instance.get<T>(path, options);

      return response.data;
    } catch (error) {
      throw error;
    }
  }

  async post<T, D>(
    path: string,
    data?: D,
    options?: AxiosRequestConfig,
  ): Promise<T> {
    try {
      const response = await this.instance.post<T>(path, data, options);

      return response.data;
    } catch (error) {
      throw error;
    }
  }

  async put<T, D>(
    path: string,
    data?: D,
    options?: AxiosRequestConfig,
  ): Promise<T> {
    try {
      const response = await this.instance.put<T>(path, data, options);

      return response.data;
    } catch (error) {
      throw error;
    }
  }

  async delete<T>(path: string, options?: AxiosRequestConfig): Promise<T> {
    try {
      const response = await this.instance.delete<T>(path, options);

      return response.data;
    } catch (error) {
      throw error;
    }
  }
}

const apiProvider = new ApiProvider()

export default apiProvider;
