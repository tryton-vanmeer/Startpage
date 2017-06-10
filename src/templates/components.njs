{% macro panel(title, body='') %}
<div class="panel panel-primary">
    <div class="panel-heading">
        <h3 class="panel-title">{{ title }}</h3>
    </div>
    <div class="panel-body">
        {{ body | safe }}
    </div>
</div>
{% endmacro %}
