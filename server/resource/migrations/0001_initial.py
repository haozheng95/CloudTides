# Generated by Django 2.2.7 on 2019-11-24 06:20

from django.db import migrations, models


class Migration(migrations.Migration):

    initial = True

    dependencies = [
    ]

    operations = [
        migrations.CreateModel(
            name='Resource',
            fields=[
                ('id', models.AutoField(auto_created=True, primary_key=True, serialize=False, verbose_name='ID')),
                ('name', models.CharField(max_length=200)),
                ('date_added', models.DateTimeField(blank=True, null=True)),
                ('host_address', models.TextField(unique=True)),
                ('platform_type', models.CharField(choices=[('1', 'vsphere'), ('2', 'kvm'), ('3', 'hyper-v')], default='vsphere', max_length=10)),
                ('username', models.CharField(max_length=150, unique=True)),
                ('password', models.CharField(max_length=128)),
                ('status', models.CharField(choices=[('1', 'idle'), ('2', 'busy'), ('3', 'contributing')], default='idle', max_length=20)),
                ('total_disk', models.FloatField(blank=True, null=True)),
                ('total_ram', models.FloatField(blank=True, null=True)),
                ('total_cpu', models.FloatField(blank=True, null=True)),
                ('current_disk', models.FloatField(blank=True, null=True)),
                ('current_ram', models.FloatField(blank=True, null=True)),
                ('current_cpu', models.FloatField(blank=True, null=True)),
                ('is_active', models.BooleanField(blank=True, null=True)),
            ],
            options={
                'db_table': 'resources',
            },
        ),
    ]
