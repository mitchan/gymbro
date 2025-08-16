class ApiClient {
  async fetch(input: RequestInfo, init?: RequestInit): Promise<unknown> {
    const response = await fetch(input, {
      ...(init ?? {}),
      credentials: "include",
    });

    if (!response.ok) {
      throw new Error("http error");
    }

    return response.json();
  }
}

export const apiClient = new ApiClient();
