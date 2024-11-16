import { http } from "@/utils/request";
import type { UserInfo, CommonResponse } from "@/types";

export const userApi = {
  // 更新用户信息
  async updateUserInfo(
    data: Partial<UserInfo>
  ): Promise<CommonResponse<UserInfo>> {
    return http.post("/user/info", data);
  },
};
