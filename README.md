
### 方法二：使用 Markdown 列表扩展（推荐，显示更清晰）

```markdown
## 📁 项目目录结构

- `gin-example/`
  - `conf/`
    - `app.ini`
  - `main.go`
  - `middleware/`
    - `jwt/`
      - `jwt.go`
  - `models/`
    - `article.go`
    - `auth.go`
    - `models.go`
    - `tag.go`
  - `pkg/`
    - `e/`
      - `code.go`
      - `msg.go`
    - `logging/`
      - `file.go`
      - `log.go`
    - `setting/`
      - `setting.go`
    - `util/`
      - `jwt.go`
      - `pagination.go`
  - `routers/`
    - `api/`
      - `auth.go`
      - `v1/`
        - `article.go`
        - `tag.go`
    - `router.go`
  - `runtime/`
