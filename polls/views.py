from django.shortcuts import render
from django.http import HttpResponse

def index(request):
  return HttpResponse("Hello, world. Welcome to the polls. Remember to vote!")

# Create your views here.
