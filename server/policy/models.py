from django.db import models
from resource.models import *

# Create your models here.
class Policy(models.Model):

    DEPLOY_TYPE = (
        ('1', 'kubernetes'),
        ('2', 'vm')
    )

    host_address = models.TextField()
    host_name = models.TextField(unique=True)
    name = models.CharField(unique=True, max_length=150)
    date_created = models.DateTimeField(blank=True, null=True)
    is_destroy = models.BooleanField(blank=True, null=True, default=False)
    deploy_type = models.CharField(max_length=20, choices=DEPLOY_TYPE, default='vm')
    idle_policy = models.TextField(blank=True, null=True)
    resource = models.ForeignKey(Resource, on_delete=models.CASCADE, null=True)

    class Meta:
        verbose_name = 'Tides Policy'
        verbose_name_plural = 'Tides Policies'

    def save(self, *args, **kwargs):
        # do something
        super().save(*args, **kwargs)
        # do something