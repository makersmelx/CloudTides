# Generated by Django 2.2.7 on 2019-12-03 14:38

from django.db import migrations, models


class Migration(migrations.Migration):

    dependencies = [
        ('resource', '0005_auto_20191201_1336'),
    ]

    operations = [
        migrations.AddField(
            model_name='resource',
            name='polling_interval',
            field=models.IntegerField(blank=True, null=True),
        ),
        migrations.AlterField(
            model_name='resource',
            name='status',
            field=models.CharField(choices=[('1', 'idle'), ('2', 'busy'), ('3', 'contributing')], default='busy', max_length=20),
        ),
    ]