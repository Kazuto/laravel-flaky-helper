# 🧪 Laravel Flaky Test Helper

A simple CLI tool to repeatedly run Laravel tests until they fail — perfect for debugging intermittent, flaky tests in your Laravel application.

---

## 🚀 Installation

### 📦 Download Prebuilt Binary

Download and extract the latest [release](https://github.com/Kazuto/laravel-flaky-helper/releases).

Then move it into a directory in your `$PATH`:

```bash
mv flaky /usr/local/bin
```

### 🛠️ Build from Source

```bash
git clone https://github.com/Kazuto/laravel-flaky-helper.git
cd laravel-flaky-helper

go build -o flaky
mv flaky /usr/local/bin
```

---

## ⚙️ Usage

Run a test repeatedly until it fails:

```bash
flaky ExampleTest
```

### ✅ Options

| Flag      | Description                                                 |
| --------- | ----------------------------------------------------------- |
| `--max N` | Maximum number of runs before stopping (default: unlimited) |

#### Example:

```bash
flaky --max 10 ExampleTest
```

<img width="516" height="273" alt="image" src="https://github.com/user-attachments/assets/3da21489-31d5-4091-abb0-0a875ff7d9c1" />


---

## 📋 Requirements

* PHP 8.0+
* Laravel 9+
* PHPUnit installed in your Laravel project

---

## 🧠 How It Works

This tool runs:

```bash
php artisan test --filter=YourTestName
```

...in a loop until it either:

* Fails (and stops immediately), or
* Reaches the optional `--max` limit

It helps expose flaky tests that sometimes pass and sometimes fail depending on runtime order, data conditions, or environment.

---

## 🤝 Contributing

Contributions, feedback, and pull requests are welcome!
Please open an [issue](https://github.com/Kazuto/laravel-flaky-helper/issues) or submit a PR.

---

## 📄 License

MIT License
© [Kazuto](https://github.com/Kazuto)
