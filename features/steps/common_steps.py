# -*- coding: utf-8 -*-
import sys
#reload(sys)
#sys.setdefaultencoding('gb2312')

import json
import time
import shutil
import os
from datetime import datetime, timedelta
import subprocess

from behave import *
from features.bdd import client as bdd_client
from features.bdd import util as bdd_util


@Given(u"{user}登录系统")
def step_impl(context, user):
	context.client = bdd_client.login('backend', user, password=None, context=context)

@When(u"{user}登录系统")
def step_impl(context, user):
	context.client = bdd_client.login('backend', user, password=None, context=context)

@Given(u"{user}注册为App用户")
def step_impl(context, user):
	context.client = bdd_client.login('app', user, password=None, context=context)
	resp = context.client.post('gskep:account.user', {
		'name': user,
		'thumbnail': 'http://resource.vxiaocheng.com/veeno/demo/girls/%s/avatar.jpg' % user
	})
	bdd_util.assert_api_call_success(resp)

@Given(u"{user}登录App")
def step_impl(context, user):
	context.client = bdd_client.login('app', user, password=None, context=context)
	context.is_app_user = True

@Given(u"{user}访问'{corpuser_name}'的商城")
def step_impl(context, user, corpuser_name):
	from features.bdd.client import RestClient
	client = RestClient()

	corp_id =bdd_util.get_corp_id_for_corpuser(client, corpuser_name)

	resp = client.put('gskep:account.mp_user', {
		'uid': user,
		'unionid': user,
		'name': user,
		'avatar': 'http://resource.vxiaocheng.com/veeno/demo/girls/%s/avatar.jpg' % user,
		'sex': 'female',
		'province': '',
		'is_subscribed': 'true',
		'corp_id': corp_id
	})
	bdd_util.assert_api_call_success(resp)
	client.jwt_token = resp.data['sid']
	client.cur_user_id = resp.data['id']
	#context.corp_token = bdd_util.get_corp_token_for_corpuser(client, corpuser_name)
	context.corpuser_name = corpuser_name
	context.corp_id = corp_id
	context.client = client

	corp_token = bdd_util.get_corp_token_for_corpuser(client, corpuser_name)
	context.client.add_cookie("__cs", corp_token)

@given(u"重置服务")
def step_impl(context):
	from features.bdd.client import RestClient
	rest_client = RestClient()
	response = rest_client.put('ginger-account:dev.bdd_reset')
	bdd_util.assert_api_call_success(response)

@then(u"结束测试")
def step_impl(context):
	import sys
	sys.exit(1)