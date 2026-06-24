export async function customFetch<T>(url: string, options: RequestInit): Promise<T> {
  const response = await fetch(url, {
    ...options,
    credentials: 'include',
  });

  const body = [204, 205, 304].includes(response.status) ? undefined : await response.text();
  const data = body ? JSON.parse(body) : {};

  return { status: response.status, data, headers: response.headers } as T;
}
