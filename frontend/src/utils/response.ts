/**
 * Check if an API response is successful (status < 400)
 * @param response - The API response object with status property
 * @returns true if response status is successful, false otherwise
 */
export function isOk(response: { status: number }): boolean {
  return response.status < 400;
}

/**
 * Type guard to check if an API response is an error (not OK and contains error string in data)
 * Narrows the type so that after this check, response.data is guaranteed to NOT be a string
 * @param response - The API response object with status and data properties
 * @returns true if response is not OK and data is a string (error message), false otherwise
 */
export function isErrorResponse<T>(response: { status: number, data: unknown }): response is { status: number, data: string } {
  return !isOk(response) && typeof response.data === 'string';
}

/**
 * Type guard to narrow response.data to a specific type T when response is successful
 * Use this to assert response.data is of type T (not a string error)
 * @param response - The API response object with status and data properties
 * @returns true if response is OK and data is not a string, false otherwise
 */
export function isSuccess<T>(response: { status: number, data: unknown }): response is { status: number, data: T } {
  return isOk(response) && typeof response.data !== 'string';
}

export function getResponseResult(response: { status: number, data: unknown }) {
  return {
    reason: response.data,
    statusCode: response.status,
  };
}
