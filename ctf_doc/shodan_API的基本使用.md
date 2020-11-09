#### shodan API search

得到所有信息（默认情况下旗标中的大字段，比如“html”会被删减）

```python
results = api.search('has_screenshot:true', minify=False)
```

默认情况下只能获得第1页数据，取得其他页数据需要用“page”参数

```python
results = api.search('apache', page=2)
```

获取shodan所有结果

```python
for banner in api.search_cursor('Server: Netwave IP Camera'):
    print(banner['ip_str'])
```

