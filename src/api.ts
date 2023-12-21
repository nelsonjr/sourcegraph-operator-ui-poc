export const adminPassword: { password: string | undefined } = { password: "" };

export const call = (
  input: RequestInfo | URL,
  init?: RequestInit
): Promise<Response> => {
  return fetch(input, {
    ...init,
    headers: {
      ...init?.headers,
      "admin-password": adminPassword.password ?? "no-password",
    },
  });
};
