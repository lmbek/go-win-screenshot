#include <windows.h>

void OnButtonClick(HWND hwnd) {
    MessageBox(hwnd, "Button Clicked!", "Hello", MB_OK);
}

LRESULT CALLBACK WindowProc(HWND hwnd, UINT uMsg, WPARAM wParam, LPARAM lParam) {
    switch (uMsg) {
        case WM_COMMAND:
            if (LOWORD(wParam) == 1) {
                OnButtonClick(hwnd);
            }
            break;

        case WM_DESTROY:
            PostQuitMessage(0);
            break;

        default:
            return DefWindowProc(hwnd, uMsg, wParam, lParam);
    }
    return 0;
}

void Main() {
    WNDCLASS wc = {0};
    wc.lpfnWndProc = WindowProc;
    wc.hInstance = GetModuleHandle(NULL);
    wc.lpszClassName = "MyClass";
    RegisterClass(&wc);

    HWND hwnd = CreateWindowEx(
        0, "MyClass", "My Window",
        WS_OVERLAPPEDWINDOW, CW_USEDEFAULT, CW_USEDEFAULT,
        500, 300, NULL, NULL, GetModuleHandle(NULL), NULL);

    HWND button = CreateWindow(
        "BUTTON", "Click Me", WS_TABSTOP | WS_VISIBLE | WS_CHILD | BS_DEFPUSHBUTTON,
        10, 10, 100, 30, hwnd, (HMENU)1, GetModuleHandle(NULL), NULL);

    ShowWindow(hwnd, SW_SHOWNORMAL);

    MSG msg = {0};
    while (GetMessage(&msg, NULL, 0, 0)) {
        TranslateMessage(&msg);
        DispatchMessage(&msg);
    }
}
