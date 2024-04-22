#include <windows.h>

// Function declarations
LRESULT CALLBACK WindowProc(HWND hwnd, UINT uMsg, WPARAM wParam, LPARAM lParam);

int WINAPI WinMain(HINSTANCE hInstance, HINSTANCE hPrevInstance, LPSTR lpCmdLine, int nCmdShow) {
    const char CLASS_NAME[] = "Sample Window Class";

    WNDCLASS wc = {};
    wc.lpfnWndProc = WindowProc;
    wc.hInstance = hInstance;
    wc.lpszClassName = CLASS_NAME;
    wc.hbrBackground = (HBRUSH)(COLOR_WINDOW + 1); // Set background color

    RegisterClass(&wc);

    HWND hwnd = CreateWindowEx(
        0,
        CLASS_NAME,
        "Advanced Windows GUI Example",
        WS_OVERLAPPEDWINDOW,
        CW_USEDEFAULT, CW_USEDEFAULT, 800, 600,
        NULL,
        NULL,
        hInstance,
        NULL
    );

    if (hwnd == NULL) {
        return 0;
    }

    ShowWindow(hwnd, nCmdShow);
    UpdateWindow(hwnd);

    // Create a menu
    HMENU hMenu = CreateMenu();
    HMENU hSubMenu = CreateMenu();
    AppendMenu(hSubMenu, MF_STRING, 1, "File");
    AppendMenu(hMenu, MF_POPUP, (UINT_PTR)hSubMenu, "Menu");

    SetMenu(hwnd, hMenu);

    // Create additional buttons
    CreateWindow(
        "BUTTON", "Click me", WS_TABSTOP | WS_VISIBLE | WS_CHILD | BS_DEFPUSHBUTTON,
        10, 10, 100, 30, hwnd, NULL, hInstance, NULL
    );

    CreateWindow(
        "BUTTON", "Another Button", WS_TABSTOP | WS_VISIBLE | WS_CHILD | BS_DEFPUSHBUTTON,
        120, 10, 150, 30, hwnd, NULL, hInstance, NULL
    );

    // Create a text field
    CreateWindow(
        "EDIT", "", WS_VISIBLE | WS_CHILD | ES_MULTILINE | ES_AUTOVSCROLL | WS_BORDER,
        10, 50, 300, 200, hwnd, NULL, hInstance, NULL
    );

    MSG msg = {};
    while (GetMessage(&msg, NULL, 0, 0)) {
        TranslateMessage(&msg);
        DispatchMessage(&msg);
    }

    return 0;
}

LRESULT CALLBACK WindowProc(HWND hwnd, UINT uMsg, WPARAM wParam, LPARAM lParam) {
    switch (uMsg) {
        case WM_DESTROY:
            PostQuitMessage(0);
            return 0;

        case WM_COMMAND:
            if (LOWORD(wParam) == 1) {
                MessageBox(hwnd, "File menu clicked!", "Info", MB_OK);
            } else if (LOWORD(wParam) == 2) {
                MessageBox(hwnd, "Another Button clicked!", "Info", MB_OK);
            }
            return 0;

        default:
            return DefWindowProc(hwnd, uMsg, wParam, lParam);
    }
}
